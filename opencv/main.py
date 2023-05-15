import cv2
from matplotlib import pyplot as plt

#1
src = cv2.imread('./files/test.jpeg')
hog = cv2.HOGDescriptor()
hog.setSVMDetector(cv2.HOGDescriptor_getDefaultPeopleDetector())

#2
loc1, weights1 = hog.detect(src)
print('len(loc1)=', len(loc1))
dst1 = src.copy()
w, h = hog.winSize
for pt in loc1:
    x, y = pt
    cv2.rectangle(dst1, (x, y), (x + w, y + h), (255, 0, 0), 2)
cv2.imshow('dst1', dst1)

#3
dst2 = src.copy()
loc2, weights2 = hog.detectMultiScale(src)
print('len(loc2)=', len(loc2))
for rect in loc2:
    x, y, w, h = rect
    cv2.rectangle(dst2, (x, y), (x + w, y + h), (0, 255, 0), 2)
cv2.imshow('dst2', dst2)

#4
dst3 = src.copy()
loc3, weights3 = hog.detectMultiScale(src)
print('len(loc3)=', len(loc3))
for i, rect in enumerate(loc3):
    x, y, w, h = rect
    if weights3[i] > 0.5:
        cv2.rectangle(dst3, (x, y), (x + w, y + h), (0, 0, 255), 2)
    else:
        cv2.rectangle(dst3, (x, y), (x + w, y + h), (255, 0, 0), 2)


cv2.imshow('dst3', dst3)
cv2.waitKey()
cv2.destroyAllWindows()
