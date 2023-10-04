from src.repositories.cottages import CottageRepository
from src.schemas.cottages import CottageCreate, CottageRead


class CottageService:
    @classmethod
    async def add_cottage(cls, cottage: CottageCreate) -> CottageRead:
        new_cottage = await CottageRepository.add(cottage.model_dump())
        return new_cottage

    @classmethod
    async def get_cottages(cls) -> list[CottageRead]:
        cottages = await CottageRepository.list()
        return cottages
