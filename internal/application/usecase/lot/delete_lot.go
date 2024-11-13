package lot

import (
    "context"
    "auction-system/internal/domain/repository"
)

type DeleteLotUseCase struct {
    lotRepo repository.LotRepository
}

func NewDeleteLotUseCase(lotRepo repository.LotRepository) *DeleteLotUseCase {
    return &DeleteLotUseCase{
        lotRepo: lotRepo,
    }
}

func (uc *DeleteLotUseCase) Execute(ctx context.Context, id int64) error {
    _, err := uc.lotRepo.GetByID(ctx, id)
    if err != nil {
        return err
    }

    return uc.lotRepo.Delete(ctx, id)
}
