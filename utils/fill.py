import random
import os

russian_letters = 'абвгдеёжзийклмнопрстуфхцчшщъыьэюяЙЦУКЕНГШЩЗХЪЭЖДЛОРПАВЫФЯЧСМИТЬБЮ'
eng_letters = 'abcdefghijklmnopqrstuvwxyz'
def generate_random_word(length):
    return ''.join(random.choice(russian_letters) for _ in range(length))

def generate_text_file(file_path, target_size_mb):
    target_size_bytes = target_size_mb * 1024 * 1024
    current_size = os.path.getsize(file_path) if os.path.exists(file_path) else 0

    with open(file_path, 'w') as file:
        while current_size < target_size_bytes:
            word = generate_random_word(random.randint(5, 10))  # Adjust length of random word as needed
            file.write(word + ' ')
            current_size = os.path.getsize(file_path)

    print(f"File '{file_path}' has reached the target size of {target_size_mb} MB.")

def generate_repeating_sequences(file_path, target_size_mb):
    target_size_bytes = target_size_mb * 1024 * 1024
    current_size = os.path.getsize(file_path) if os.path.exists(file_path) else 0

    with open(file_path, 'w') as file:
        while current_size < target_size_bytes:
            sequence_length = random.randint(5, 15)
            character = random.choice(russian_letters)
            sequence = character * sequence_length
            file.write(sequence + ' ')
            current_size = os.path.getsize(file_path)

    print(f"File '{file_path}' has reached the target size of {target_size_mb} MB.")

# Fill test .txt files
if __name__ == "__main__":
    file_path_1 = "test_data/txt/random.txt"
    file_path_2 = "test_data/txt/sequences.txt"
    target_size_mb = 1

    generate_text_file(file_path_1, target_size_mb)
    generate_repeating_sequences(file_path_2, target_size_mb)