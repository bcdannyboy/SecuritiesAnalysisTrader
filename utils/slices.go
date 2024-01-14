package utils

func StringSliceContains(slice []string, value string) bool {
	/*
	 * Checks if the specified string slice contains the specified value.
	 *
	 * Parameters:
	 * 	slice ([]string): The slice to check.
	 * 	value (string): The value to check for.
	 *
	 * Returns:
	 * 	bool: Whether or not the slice contains the value.
	 */

	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}
