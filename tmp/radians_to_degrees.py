from math import pi

def radians_to_degrees(radians):
    degrees = radians * (180 / pi)
    return round(degrees, 1)

def main():
    radians = float(input("Enter the angle in radians: "))
    degrees = radians_to_degrees(radians)
    print(f"{radians} radians is equal to {degrees} degrees.")

if __name__ == "__main__":
    main()
    