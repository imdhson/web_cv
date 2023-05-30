import cv2
import sys

# 이미지를 불러옵니다.
# path = sys.argv[1]
path = '../files/ive.png'
img = cv2.imread(path)


# Super Resolution을 하기 위해 2배 확대를 위한 ESPCN 모델을 사용합니다.
sr = cv2.dnn_superres.DnnSuperResImpl_create()
# 모델 파일의 이름
sr.readModel('EDSR_x2.pb')
# 모델 파일에 적힌 숫자와 일치하도록 해야 합니다.
sr.setModel("edsr",2)
# img를 입력 받아 결과로 result를 돌려줍니다.  
result = sr.upsample(img)

cv2.imwrite(path, result)