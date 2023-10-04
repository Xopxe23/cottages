from sqlalchemy import String
from sqlalchemy.orm import Mapped, mapped_column

from src.database import Base
from src.schemas.cottages import CottageRead


class Cottage(Base):
    __tablename__ = 'cottages'

    id: Mapped[int] = mapped_column(primary_key=True)
    address: Mapped[str] = mapped_column(String(255), nullable=False)
    city: Mapped[str] = mapped_column(String(50), nullable=False)

    def to_read_model(self) -> CottageRead:
        return CottageRead(
            id=self.id,
            address=self.address,
            city=self.city
        )
