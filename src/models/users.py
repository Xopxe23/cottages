from fastapi import Depends
from fastapi_users_db_sqlalchemy import SQLAlchemyUserDatabase, SQLAlchemyBaseUserTable, SQLAlchemyBaseOAuthAccountTable
from sqlalchemy import Integer, ForeignKey
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import Mapped, relationship, mapped_column, declared_attr

from src.database import Base, get_async_session


class OAuthAccount(SQLAlchemyBaseOAuthAccountTable[int], Base):
    id: Mapped[int] = mapped_column(Integer, primary_key=True)

    @declared_attr
    def user_id(cls) -> Mapped[int]:
        return mapped_column(Integer, ForeignKey("user.id", ondelete="cascade"), nullable=False)


class User(SQLAlchemyBaseUserTable[int], Base):
    id: Mapped[int] = mapped_column(primary_key=True)
    oauth_accounts: Mapped[list[OAuthAccount]] = relationship(
        "OAuthAccount", lazy="joined"
    )


async def get_user_db(session: AsyncSession = Depends(get_async_session)):
    yield SQLAlchemyUserDatabase(session, User, OAuthAccount)
