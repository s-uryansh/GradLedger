import cv2
import numpy as np
from PIL import Image, ImageEnhance, ImageOps
from io import BytesIO
from typing import Tuple
from config import config
from exceptions import ImageProcessingException

def validate_and_preprocess_image(image_bytes: bytes) -> np.ndarray:
    """Validate and preprocess image for face recognition"""
    try:
        # Validate file size
        if len(image_bytes) > config.MAX_FILE_SIZE:
            raise ImageProcessingException("File size exceeds maximum allowed (10MB)")
        
        # Load and validate image
        try:
            pil_image = Image.open(BytesIO(image_bytes))
            pil_image.verify()  # Verify it's a valid image
            
            # Reload image (verify() closes the file)
            pil_image = Image.open(BytesIO(image_bytes))
            image_array = np.array(pil_image)
        except Exception:
            raise ImageProcessingException("Invalid image format or corrupted file")
        
        # Convert to RGB if needed
        if image_array.ndim == 2:  # Grayscale
            image_array = cv2.cvtColor(image_array, cv2.COLOR_GRAY2RGB)
        elif len(image_array.shape) == 3 and image_array.shape[2] == 4:  # RGBA
            image_array = cv2.cvtColor(image_array, cv2.COLOR_RGBA2RGB)
        elif len(image_array.shape) == 3 and image_array.shape[2] == 3:
            # Ensure it's RGB (PIL loads as RGB, but just to be safe)
            pass
        else:
            raise ImageProcessingException("Unsupported image format")
        
        # Resize if too large
        height, width = image_array.shape[:2]
        if max(height, width) > config.MAX_IMAGE_SIZE:
            if width > height:
                scale = config.MAX_IMAGE_SIZE / width
            else:
                scale = config.MAX_IMAGE_SIZE / height
            
            new_width = int(width * scale)
            new_height = int(height * scale)
            
            image_array = cv2.resize(
                image_array, 
                (new_width, new_height), 
                interpolation=cv2.INTER_AREA
            )
        
        # Enhance image quality
        image_array = enhance_image_quality(image_array)
        
        return image_array
        
    except (ImageProcessingException):
        raise
    except Exception as e:
        raise ImageProcessingException(f"Image preprocessing failed: {str(e)}")

def enhance_image_quality(image_array: np.ndarray) -> np.ndarray:
    """Enhance image quality for better face recognition"""
    try:
        # Convert to PIL for enhancement
        pil_image = Image.fromarray(image_array)
        
        # Auto-contrast
        pil_image = ImageOps.autocontrast(pil_image, cutoff=1)
        
        # Slight contrast enhancement
        enhancer = ImageEnhance.Contrast(pil_image)
        pil_image = enhancer.enhance(1.1)
        
        # Slight sharpness enhancement
        enhancer = ImageEnhance.Sharpness(pil_image)
        pil_image = enhancer.enhance(1.05)
        
        return np.array(pil_image)
    except Exception:
        # Return original if enhancement fails
        return image_array

