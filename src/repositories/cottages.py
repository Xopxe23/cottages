from sqlalchemy import insert, select

from src.database import async_session_maker
from src.models.cottages import Cottage
from src.schemas.cottages import CottageRead


class CottageRepository:
    @classmethod
    async def add(cls, data: dict) -> CottageRead:
        query = insert(Cottage).values(**data).returning(Cottage)
        async with async_session_maker() as session:
            new_cottage = await session.execute(query)
            await session.commit()
        result = new_cottage.scalar_one().to_read_model()
        return result

    @classmethod
    async def list(cls) -> list[CottageRead]:
        query = select(Cottage)
        async with async_session_maker() as session:
            cottages = await session.execute(query)
        result = [row[0].to_read_model() for row in cottages.all()]
        return result
