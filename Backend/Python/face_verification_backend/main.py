from fastapi import FastAPI, UploadFile, File, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse
from models import VerificationResponse
from face_service import FaceVerificationService
from exceptions import NoFaceDetectedException, ImageProcessingException
import warnings

warnings.filterwarnings("ignore", category=UserWarning)

app = FastAPI(title="Face Verification Service", version="1.0.0")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
face_service = FaceVerificationService()

@app.post("/verify", response_model=VerificationResponse)
async def verify_faces(
    image1: UploadFile = File(..., description="Live captured image"),
    image2: UploadFile = File(..., description="Reference image (ID card)")
):
    """
    Verify if faces in two images match.
    Returns detailed verification results with confidence scores.
    """
    try:
        allowed_types = {"image/jpeg", "image/png", "image/jpg", "image/webp"}
        
        if image1.content_type not in allowed_types:
            raise HTTPException(
                status_code=400, 
                detail="Invalid file type for image1. Allowed: JPEG, PNG, WebP"
            )
        
        if image2.content_type not in allowed_types:
            raise HTTPException(
                status_code=400, 
                detail="Invalid file type for image2. Allowed: JPEG, PNG, WebP"
            )
        
        image1_bytes = await image1.read()
        image2_bytes = await image2.read()
        
        result = await face_service.verify_faces(image1_bytes, image2_bytes)
        
        return VerificationResponse(**result)
        
    except (NoFaceDetectedException, ImageProcessingException) as e:
        raise e
    except HTTPException as e:
        raise e
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Internal server error: {str(e)}")

@app.get("/health")
async def health_check():
    """Simple health check endpoint"""
    return {"status": "healthy", "service": "Face Verification API"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="127.0.0.1", port=8000, reload=True)