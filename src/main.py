import uvicorn
from fastapi import FastAPI

from src.api.cottages import router as cottages_router
from src.api.users import router as users_router

app = FastAPI(title="Bookings")
app.include_router(router=cottages_router)
app.include_router(router=users_router)

if __name__ == "__main__":
    uvicorn.run("src/main:app", reload=True)
