export function Header() {
  return (
    <header className="border-b border-gray-200 dark:border-gray-800">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="text-center">
          <h1 className="text-4xl font-bold text-gray-900 dark:text-gray-100 mb-2">
            HookHub
          </h1>
          <p className="text-lg text-gray-600 dark:text-gray-400 mb-4">
            Discover Claude Code Hooks
          </p>
          <p className="text-sm text-gray-500 dark:text-gray-500 max-w-2xl mx-auto">
            Browse a curated collection of open source hooks to enhance your Claude Code workflow.
            Automate formatting, validation, notifications, and more.
          </p>
        </div>
      </div>
    </header>
  );
}
