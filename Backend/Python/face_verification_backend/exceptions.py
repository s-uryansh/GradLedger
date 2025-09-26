from fastapi import HTTPException

class NoFaceDetectedException(HTTPException):
    def __init__(self, detail: str = "No face detected in image"):
        super().__init__(status_code=400, detail=detail)

class ImageProcessingException(HTTPException):
    def __init__(self, detail: str = "Error processing image"):
        super().__init__(status_code=422, detail=detail)
