from fastapi import APIRouter

from src.schemas.cottages import CottageCreate, CottageRead
from src.services.cottages import CottageService

router = APIRouter(prefix="/cottages", tags=["Cottages"])


@router.post("/")
async def create_cottage(data: CottageCreate) -> CottageRead:
    cottage = await CottageService.add_cottage(data)
    return cottage


@router.get("/")
async def get_cottages() -> list[CottageRead]:
    cottages = await CottageService.get_cottages()
    return cottages
