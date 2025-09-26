class Config:
    # Face Recognition Settings
    # CHANGE THIS VALUE from 0.6 to 0.55 for better accuracy
    FACE_MATCH_THRESHOLD = 0.55  # <-- Stricter and more secure threshold
    MIN_FACE_SIZE = 50
    MAX_IMAGE_SIZE = 2048
    MAX_FILE_SIZE = 10 * 1024 * 1024  # 10MB

config = Config()