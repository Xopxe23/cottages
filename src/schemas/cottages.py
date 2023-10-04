from pydantic import BaseModel


class CottageRead(BaseModel):
    id: int
    address: str
    city: str

    class ConfigDict:
        from_attributes = True


class CottageCreate(BaseModel):
    address: str
    city: str
