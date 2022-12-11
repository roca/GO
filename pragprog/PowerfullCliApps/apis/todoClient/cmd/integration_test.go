//go:build integration

package cmd

func randomTaskName(t *testing.T) string {
	t.Helper()
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var p strings.Builder
	for i := 0; i < 32; i++ {
		p.WriteByte(chars[r.Intn(len(chars))])
	}

	return p.String()
}

func TestIntegration(t *testing.T) {
	t.Run("AddTask", func(t *testing.T) {})
	t.Run("ListTask", func(t *testing.T) {})
	t.Run("ViewTask", func(t *testing.T) {})
	t.Run("CompleteTask", func(t *testing.T) {})
	t.Run("ListCompletedTask", func(t *testing.T) {})
	t.Run("DeleteTask", func(t *testing.T) {})
	t.Run("ListDeletedTask", func(t *testing.T) {})
}
