func categorizeBox(length int, width int, height int, mass int) string {
    volume := length * width * height
    if (length >= 10e3 || width >= 10e3 || height>= 10e3 || volume >= 10e8) && mass >= 100 {
        return "Both"
    } else if mass>= 100 && volume < 10e8 {
        return "Heavy"
    } else if (length >= 10e3 || width >= 10e3 || height>= 10e3 || volume >= 10e8) && mass < 100 {
        return "Bulky"
    }
    return "Neither"
}
