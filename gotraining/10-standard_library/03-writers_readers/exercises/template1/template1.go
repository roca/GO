		return
	}
	defer file.Close()

	// Send the document to the file.
	writers = append(writers, file)

	// MultiWriter(io.Writer...) returns a single writer which multiplexes its
	// writes across all of the writers we pass in.
	dest := io.MultiWriter(writers...)

	// Write to dest the same way as before, copying from the Body.
	_, err = io.Copy(dest, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
}