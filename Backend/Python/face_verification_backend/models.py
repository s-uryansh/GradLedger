from pydantic import BaseModel
from typing import Optional, Dict, Any

class VerificationResponse(BaseModel):
    match: bool
    confidence: float
    distance: float
    message: str
    processing_time_ms: float
    metadata: Optional[Dict[str, Any]] = None

