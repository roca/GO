import { Hook } from '@/lib/types/hook';
import { CategoryBadge } from './CategoryBadge';

interface HookCardProps {
  hook: Hook;
}

export function HookCard({ hook }: HookCardProps) {
  return (
    <a
      href={hook.repoUrl}
      target="_blank"
      rel="noopener noreferrer"
      className="block p-6 bg-white dark:bg-neutral-900 border border-gray-200 dark:border-gray-800 rounded-lg hover:shadow-md transition-shadow duration-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-neutral-950"
    >
      <article>
        <div className="flex items-start justify-between gap-2 mb-3">
          <h2 className="text-lg font-semibold text-gray-900 dark:text-gray-100">
            {hook.name}
          </h2>
          <CategoryBadge category={hook.category} />
        </div>

        <div className="mb-3">
          <span className="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-gray-100 text-gray-700 dark:bg-gray-800 dark:text-gray-300">
            {hook.hookEvent}
          </span>
        </div>

        <p className="text-sm text-gray-600 dark:text-gray-400 line-clamp-3 mb-4">
          {hook.description}
        </p>

        <div className="flex items-center justify-between text-xs text-gray-500 dark:text-gray-500">
          {hook.author && (
            <span className="flex items-center gap-1">
              <svg
                className="w-4 h-4"
                fill="currentColor"
                viewBox="0 0 24 24"
                aria-hidden="true"
              >
                <path
                  fillRule="evenodd"
                  d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z"
                  clipRule="evenodd"
                />
              </svg>
              {hook.author}
            </span>
          )}
          <span className="flex items-center gap-1">
            <svg
              className="w-4 h-4"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              aria-hidden="true"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
              />
            </svg>
            View Repo
          </span>
        </div>
      </article>
    </a>
  );
}
