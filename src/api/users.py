from fastapi import APIRouter, Depends
from fastapi_users import FastAPIUsers

from src.config import settings
from src.models.users import User
from src.schemas.users import UserCreate, UserRead
from src.services.users import get_user_manager
from src.utils.auth import auth_backend, google_oauth_client

fastapi_users = FastAPIUsers[User, int](
    get_user_manager,
    [auth_backend],
)


router = APIRouter(
    prefix="/auth",
    tags=["Auth"]
)

router.include_router(
    fastapi_users.get_register_router(UserRead, UserCreate),
    prefix="",
)

router.include_router(
    fastapi_users.get_auth_router(auth_backend),
    prefix="/jwt",
)

router.include_router(
    fastapi_users.get_verify_router(UserRead),
    prefix="",
)

router.include_router(
    fastapi_users.get_oauth_router(
        google_oauth_client,
        auth_backend,
        settings.SECRET,
        is_verified_by_default=True,
    ),
    prefix="/google",
)

current_active_user = fastapi_users.current_user(active=True)


@router.get("/authenticated-route")
async def authenticated_route(user: User = Depends(current_active_user)):
    return {"message": f"Hello {user.email}!"}
