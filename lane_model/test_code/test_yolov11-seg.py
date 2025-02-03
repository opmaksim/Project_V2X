import cv2
from ultralytics import YOLO

# ✅ 모델 경로 설정
model_path = "/home/ubuntu/workspace/roboflow/Yolov11n-seg-finetuned.pt"
model = YOLO(model_path)

# ✅ 비디오 경로 설정
video_path = "/home/ubuntu/workspace/road2.mp4"
output_path = "/home/ubuntu/workspace/road2_output.mp4"

# 비디오 캡처
cap = cv2.VideoCapture(video_path)
frame_width = int(cap.get(3))
frame_height = int(cap.get(4))
fps = int(cap.get(cv2.CAP_PROP_FPS))

# 비디오 저장 설정
fourcc = cv2.VideoWriter_fourcc(*'mp4v')  # MP4 형식
out = cv2.VideoWriter(output_path, fourcc, fps, (frame_width, frame_height))

# ✅ Confidence Threshold 설정 (조절 가능)
CONF_THRESHOLD = 0.65  # (기본값 0.25 → 0.5로 설정하여 신뢰도 높은 탐지만 수행)

# 프레임별 YOLO 세그멘테이션 수행
while cap.isOpened():
    ret, frame = cap.read()
    if not ret:
        break  # 비디오 끝

    # YOLO 추론 수행 (conf 설정 추가)
    results = model(frame, conf=CONF_THRESHOLD)  # ✅ conf 값 조절 가능

    # 결과를 프레임에 그리기
    for result in results:
        frame = result.plot()  # YOLO가 자동으로 세그멘테이션을 시각화

    # 비디오 저장
    out.write(frame)

    # 화면에 출력 (선택 사항)
    cv2.imshow("YOLOv11 Segmentation", frame)
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

# 리소스 해제
cap.release()
out.release()
cv2.destroyAllWindows()

print(f"✅ 결과 저장 완료: {output_path}")
