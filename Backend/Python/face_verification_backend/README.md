# Face recognization

This is first step verificaiton step for project GradLedger.

## Working:
* Takes in 2 images, validates the file type, size and format.
* Both the images are resized, enhanced and normalized
* Both the images are processed to find and extract the faces
* Gives a  quality score to each face
* From each image highest quality face is selected
* Face encodings are compared using Euclidean distance
* Final confidence score is calculated based on distance and quality

---
#### Face Encoding
Each face encoding is 128-dimensional vector.
These numbers represent specific face features such as eye spacing, nose shape, jawline etc.
`face_recognization` library uses a deep neural network to extract these features.

Euclidean distance measures how "far apart" two points are in multi-dimensional space. It's like measuring the straight-line distance between two points.

`Distance = √[(A₁-B₁)² + (A₂-B₂)² + ... + (Aₙ-Bₙ)²]`

In `face_service.py`:
- `face_recognition.face_distance()` takes two face encoding of 128-dimensional vector.
- Computes Euclidean distance b/w them and returns a number
- Typically this number is b/w 0.0 and 1.2

### Verification of faces:
If number returned by `face_recognition.face_distance()` 
- ≤ 0.55 faces are considered a match.
- \> 0.55 faces are different.

This value is a result of hit and trial method

## Why Euclidean Distance:
Photos of different people will have vectors that are far apart but If two photos are of the same person, their 128-dimensional vectors will be close to each other in this mathematical space.

Euclidean distance provides a reliable, consistent way to measure similarity across all 128 facial feature dimensions

# Installation & Run
```bash
pip install -e . 
pip install -r requirements.txt
python3 run.py
```