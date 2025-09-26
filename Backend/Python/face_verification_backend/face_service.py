import face_recognition
import numpy as np
from typing import Tuple, List, Dict, Any
import time
from config import config
from exceptions import NoFaceDetectedException
from image_utils import validate_and_preprocess_image

class FaceVerificationService:
    def __init__(self):
        self.threshold = config.FACE_MATCH_THRESHOLD
        self.min_face_size = config.MIN_FACE_SIZE
    
    def detect_and_encode_faces(self, image_array: np.ndarray) -> Tuple[List, List, List[float]]:
        """Detect faces and return locations, encodings, and quality scores"""
        face_locations = face_recognition.face_locations(image_array, model="hog")
        
        if not face_locations:
            face_locations = face_recognition.face_locations(image_array, model="cnn")
        
        if not face_locations:
            return [], [], []
        
        valid_locations = []
        quality_scores = []
        
        for (top, right, bottom, left) in face_locations:
            face_width = right - left
            face_height = bottom - top
            
            if face_width >= self.min_face_size and face_height >= self.min_face_size:
                valid_locations.append((top, right, bottom, left))
                
                quality = self._calculate_face_quality((top, right, bottom, left), image_array.shape)
                quality_scores.append(quality)
        
        if not valid_locations:
            return [], [], []
        
        face_encodings = face_recognition.face_encodings(image_array, valid_locations)
        
        return valid_locations, face_encodings, quality_scores
    
    def _calculate_face_quality(self, face_location: Tuple, image_shape: Tuple) -> float:
        """Calculate face quality score based on size and position"""
        top, right, bottom, left = face_location
        image_height, image_width = image_shape[:2]
        
        face_area = (bottom - top) * (right - left)
        image_area = image_height * image_width
        face_ratio = face_area / image_area
        
        if 0.05 <= face_ratio <= 0.3:
            size_score = 1.0
        elif face_ratio < 0.01:
            size_score = 0.2
        elif face_ratio > 0.5:
            size_score = 0.6
        else:
            size_score = max(0.4, min(1.0, face_ratio * 5))
        
        face_center_x = (left + right) / 2
        face_center_y = (top + bottom) / 2
        image_center_x = image_width / 2
        image_center_y = image_height / 2
        
        center_distance = np.sqrt(
            ((face_center_x - image_center_x) / image_width) ** 2 +
            ((face_center_y - image_center_y) / image_height) ** 2
        )
        position_score = max(0.5, 1.0 - center_distance)
        
        return (size_score * 0.7) + (position_score * 0.3)
    
    def get_best_face_encoding(self, image_array: np.ndarray) -> Tuple[np.ndarray, float, Dict]:
        """Get the best quality face encoding from image"""
        locations, encodings, qualities = self.detect_and_encode_faces(image_array)
        
        if not encodings:
            raise NoFaceDetectedException("No suitable face detected in image")
        
        best_idx = np.argmax(qualities)
        
        return (
            encodings[best_idx], 
            qualities[best_idx],
            {
                "faces_detected": len(encodings),
                "selected_face_quality": qualities[best_idx],
                "face_location": locations[best_idx]
            }
        )
    
    def calculate_match_confidence(self, distance: float, quality1: float, quality2: float) -> Tuple[float, bool]:
        """Calculate final confidence and match decision with a more intuitive score."""
        is_match = distance <= self.threshold

        if not is_match:
            return 0.0, False
        confidence = (1.0 - (distance / self.threshold)) ** 2
        
        quality_factor = (quality1 + quality2) / 2
        
        quality_adjustment = (quality_factor - 0.5) * 0.2
        
        final_confidence = confidence + quality_adjustment
        
        final_confidence = max(0.0, min(1.0, final_confidence))
        
        return final_confidence, is_match

    async def verify_faces(self, image1_bytes: bytes, image2_bytes: bytes) -> Dict[str, Any]:
        """Main verification method"""
        start_time = time.time()
        
        try:
            img1 = validate_and_preprocess_image(image1_bytes)
            img2 = validate_and_preprocess_image(image2_bytes)
            
            encoding1, quality1, metadata1 = self.get_best_face_encoding(img1)
            encoding2, quality2, metadata2 = self.get_best_face_encoding(img2)
            
            distance = face_recognition.face_distance([encoding1], encoding2)[0]
            confidence, is_match = self.calculate_match_confidence(distance, quality1, quality2)
            
            processing_time = (time.time() - start_time) * 1000
            
            if is_match:
                message = "Faces match"
            elif distance <= self.threshold + 0.05: # A small buffer for "similar" faces
                message = "Faces are similar but do not meet the confidence threshold"
            else:
                message = "Faces do not match"
            
            if is_match:
                if confidence > 0.7:
                    message = "Faces match with high confidence"
                elif confidence > 0.5:
                    message = "Faces match with moderate confidence"
                else:
                    message = "Faces match with low confidence"

            return {
                "match": is_match,
                "confidence": round(float(confidence), 3),
                "distance": round(float(distance), 3),
                "message": message,
                "processing_time_ms": round(processing_time, 2),
                "metadata": {
                    "image1_quality": round(quality1, 3),
                    "image2_quality": round(quality2, 3),
                    "threshold_used": self.threshold,
                    "image1_faces": metadata1["faces_detected"],
                    "image2_faces": metadata2["faces_detected"]
                }
            }
            
        except Exception as e:
            raise e