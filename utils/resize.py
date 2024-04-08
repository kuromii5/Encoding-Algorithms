from PIL import Image
import os

def resize_images(input_folder, output_folder, target_size=(400, 400)):
    if not os.path.exists(output_folder):
        os.makedirs(output_folder)

    for filename in os.listdir(input_folder):
        if filename.endswith(('.png', '.jpg', '.jpeg', '.gif')):
            with Image.open(os.path.join(input_folder, filename)) as img:
                img_resized = img.resize(target_size, Image.LANCZOS)
                img_resized.save(os.path.join(output_folder, filename))

    print("All images resized successfully.")

# Resize every image in folder to equal pixel size
if __name__ == "__main__":
    input_folder = "test_data/images/input_images"
    output_folder = "test_data/images/resized_images"
    target_size = (400, 400)
    resize_images(input_folder, output_folder, target_size)