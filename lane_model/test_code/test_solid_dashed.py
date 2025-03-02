import cv2
import time
import numpy as np
from ultralytics import YOLO

# ✅ 모델 및 비디오 경로 설정
model_path = "/home/ubuntu/workspace/roboflow/Yolov11n-seg-finetuned.pt"
video_path = "/home/ubuntu/workspace/road.mp4"
output_path = "/home/ubuntu/workspace/road_output.mp4"

# ✅ YOLO 모델 로드 (GPU 사용)
model = YOLO(model_path).to('cuda')

# ✅ 비디오 캡처 및 저장 설정
cap = cv2.VideoCapture(video_path)
frame_width, frame_height = int(cap.get(3)), int(cap.get(4))
fps = int(cap.get(cv2.CAP_PROP_FPS))
out = cv2.VideoWriter(output_path, cv2.VideoWriter_fourcc(*'mp4v'), fps, (frame_width, frame_height))

# ✅ Confidence Threshold 설정
CONF_THRESHOLD = 0.6  # 신뢰도 60% 이상 탐지만 수행

# ✅ 클래스 정의 (data.yaml 기반)
CLASS_NAMES = ['Lane-Markings', 'Left Boundary -Dashed-', 'Left Boundary -Solid-', 
               'Right Boundary -Dashed-', 'Right Boundary -Solid-']

# ✅ Solid & Dashed 클래스 ID 매핑
SOLID_CLASS_IDS = [CLASS_NAMES.index('Left Boundary -Solid-'), CLASS_NAMES.index('Right Boundary -Solid-')]
DASHED_CLASS_IDS = [CLASS_NAMES.index('Left Boundary -Dashed-'), CLASS_NAMES.index('Right Boundary -Dashed-')]

print(f"Solid Class IDs: {SOLID_CLASS_IDS}, Dashed Class IDs: {DASHED_CLASS_IDS}")

# ✅ FPS 계산용 변수 초기화
start_time = time.time()
frame_count = 0

# ✅ 프레임별 YOLO 세그멘테이션 수행
while cap.isOpened():
    ret, frame = cap.read()
    if not ret:
        break  # 비디오 종료 시 반복문 탈출

    frame_count += 1
    elapsed_time = time.time() - start_time
    fps = frame_count / elapsed_time

    # ✅ YOLO 추론 수행
    results = model(frame, conf=CONF_THRESHOLD)

    # ✅ 탐지된 객체 저장
    filtered_boxes = []
    filtered_masks = []

    for result in results:
        if result.boxes is not None and result.masks is not None:
            for i, (box, cls_id) in enumerate(zip(result.boxes.xyxy, result.boxes.cls)):
                cls_id = int(cls_id)
                if cls_id in SOLID_CLASS_IDS or cls_id in DASHED_CLASS_IDS:
                    filtered_boxes.append((box, cls_id))

                    # ✅ 마스크 존재 여부 확인 후 추가
                    if result.masks is not None and i < len(result.masks.xy):
                        filtered_masks.append((result.masks.xy[i], cls_id))

    # ✅ 필터링된 결과를 프레임에 표시
    for box, cls_id in filtered_boxes:
        x1, y1, x2, y2 = map(int, box)
        color = (0, 255, 0) if cls_id in SOLID_CLASS_IDS else (0, 0, 255)  # solid=green, dashed=red
        cv2.rectangle(frame, (x1, y1), (x2, y2), color, 2)
        label = "Solid" if cls_id in SOLID_CLASS_IDS else "Dashed"
        cv2.putText(frame, label, (x1, y1 - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.5, color, 2)

    for mask, cls_id in filtered_masks:
        color = (0, 255, 0, 100) if cls_id in SOLID_CLASS_IDS else (0, 0, 255, 100)
        mask = np.array(mask, dtype=np.int32)  # 마스크를 정수형 배열로 변환
        cv2.fillPoly(frame, [mask], color[:3])  # 마스크 색상 적용

    # ✅ 실시간 FPS 표시
    cv2.putText(frame, f"FPS: {fps:.2f}", (10, 30), cv2.FONT_HERSHEY_SIMPLEX, 0.7, (255, 255, 255), 2)

    # ✅ 비디오 저장 및 출력
    out.write(frame)
    cv2.imshow("YOLOv11 Segmentation", frame)
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

# ✅ 리소스 해제
cap.release()
out.release()
cv2.destroyAllWindows()

print(f"✅ 결과 저장 완료: {output_path}")
