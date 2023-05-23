import cv2
import cvlib as cv
from IPython.display import Image, display

path = "./files/ive.png" # 사진 파일의 디렉토리
display(Image(filename = path))

img = cv2.imread(path)
conf = 0.5
model_name = "yolov3"

result = cv.detect_common_objects(img, confidence=conf, model=model_name)

