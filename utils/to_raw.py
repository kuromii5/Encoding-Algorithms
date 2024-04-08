from PIL import Image
import numpy as np
import os

def binarize_image(input_image_path, output_raw_path, threshold=128):
    img = Image.open(input_image_path)
    input_size = os.path.getsize(input_image_path)

    img_gray = img.convert('L')

    img_bin = img_gray.point(lambda p: p > threshold and 255)
    img_array = np.array(img_bin)

    img_array_binary = np.where(img_array > threshold, 1, 0)

    packed_data = np.packbits(img_array_binary)

    with open(output_raw_path, 'wb') as f:
        f.write(packed_data)

    file_size = os.path.getsize(output_raw_path)
    print(f"Image '{input_image_path}' binarized and saved as RAW format at '{output_raw_path}'.\nInput size:'{input_size}' \nNew size: '{file_size}'")

def image_to_raw(input_image_path, output_raw_path):
    img = Image.open(input_image_path)
    input_size = os.path.getsize(input_image_path)

    img_gray = img.convert('L')

    img_array = np.array(img_gray)

    with open(output_raw_path, 'wb') as f:
        f.write(img_array.tobytes())

    file_size = os.path.getsize(output_raw_path)
    print(f"Image '{input_image_path}' converted to RAW format and saved as '{output_raw_path}'.\nInput size:'{input_size}' \nNew size: '{file_size}'")

def colored_to_raw(input_image_path, output_raw_path):
    img = Image.open(input_image_path)
    input_size = os.path.getsize(input_image_path)

    img_array = np.array(img)

    with open(output_raw_path, 'wb') as f:
        f.write(img_array.tobytes())

    file_size = os.path.getsize(output_raw_path)
    print(f"Image '{input_image_path}' converted to RAW format and saved as '{output_raw_path}'.\nInput size:'{input_size}' \nNew size: '{file_size}'")

# Make RAW images
if __name__ == "__main__":
    bin_path = "test_data/images/resized_images/1.jpg"
    output_bin_path = "raw_images/1.raw"
    binarize_image(bin_path, output_bin_path)

    gray_path = "test_data/images/resized_images/2.jpg"
    output_gray_path = "raw_images/2.raw"
    image_to_raw(gray_path, output_gray_path)

    colored_path = "test_data/images/resized_images/3.jpg"
    output_colored_path = "raw_images/3.raw"
    colored_to_raw(colored_path, output_colored_path)