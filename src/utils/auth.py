from fastapi_users.authentication import (AuthenticationBackend,
                                          BearerTransport, JWTStrategy)
from httpx_oauth.clients.google import GoogleOAuth2

from src.config import settings

bearer_transport = BearerTransport(tokenUrl="auth/jwt/login")

SECRET = settings.SECRET


def get_jwt_strategy() -> JWTStrategy:
    return JWTStrategy(secret=SECRET, lifetime_seconds=3600)


auth_backend = AuthenticationBackend(
    name="jwt",
    transport=bearer_transport,
    get_strategy=get_jwt_strategy,
)

google_oauth_client = GoogleOAuth2("CLIENT_ID", "CLIENT_SECRET")
