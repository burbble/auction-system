package lot

import (
    "context"
    "auction-system/internal/application/dto/lot"
    "auction-system/internal/domain/repository"
    "auction-system/internal/domain/errors"
)

type GetLotsUseCase struct {
    lotRepo repository.LotRepository
}

func NewGetLotsUseCase(lotRepo repository.LotRepository) *GetLotsUseCase {
    return &GetLotsUseCase{
        lotRepo: lotRepo,
    }
}

func (uc *GetLotsUseCase) Execute(ctx context.Context, page, pageSize int) (*lot.LotListResponse, error) {
    if page < 1 {
        return nil, errors.New(errors.ErrorTypeValidation, "page number must be greater than 0", nil)
    }
    if pageSize < 1 {
        return nil, errors.New(errors.ErrorTypeValidation, "page size must be greater than 0", nil)
    }
    
    const maxPageSize = 100
    if pageSize > maxPageSize {
        pageSize = maxPageSize
    }

    offset := (page - 1) * pageSize

    lots, total, err := uc.lotRepo.List(ctx, offset, pageSize)
    if err != nil {
        return nil, err
    }

    if len(lots) == 0 && page > 1 {
        return nil, errors.New(errors.ErrorTypeNotFound, "no lots found for this page", nil)
    }

    return lot.ToLotListResponse(lots, total), nil
}
