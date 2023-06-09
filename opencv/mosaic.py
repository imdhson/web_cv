import cvlib as cv
import cv2
import sys

path = sys.argv[1]
# path = 'IU1.jpg' #-- 이미지 파일 경로 입력
img = cv2.imread(path)

 
face, confidence = cv.detect_face(img)

for idx, f in enumerate(face):
    (startX, startY) = f[0], f[1]
    (endX, endY) = f[2], f[3]
 
    #-- 모자이크 효과 주기
    face_region = img[startY:endY, startX:endX]
    B = face_region.shape[0]
    S = face_region.shape[1]

    face_region = cv2.resize(face_region, None, fx=0.05, fy=0.05, interpolation=cv2.INTER_AREA)
    face_region = cv2.resize(face_region, (S, B), interpolation=cv2.INTER_AREA)
    img[startY:endY, startX:endX] = face_region
 
cv2.imwrite(path, img) # 반영된 이미지로 파일 대체