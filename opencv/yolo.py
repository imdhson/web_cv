import cv2
import cvlib as cv
import sys
import numpy as np

path = sys.argv[1]

img = cv2.imread(path) # 이미지 파일 불러오기
conf = 0.5 # confidence 역치 값
model_name = "yolov3" # 사물 인식 모델 이름

result = cv.detect_common_objects(img, confidence=conf, model=model_name)

output_path = path
result_img = cv.object_detection.draw_bbox(img, *result) # result 결과를 이미지에 반영
cv2.imwrite(out_path, result_img) # 반영된 이미지로 파일 대체



