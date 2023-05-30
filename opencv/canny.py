import cv2
import sys

path = sys.argv[1]
gray_img = cv2.imread(path, cv2.IMREAD_GRAYSCALE)

threshold1 = 0
threshold2 = 360
edge_img = cv2.Canny(gray_img, threshold1, threshold2)
cv2.imwrite(path, edge_img)