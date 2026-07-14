import { HookCategory } from '@/lib/types/hook';

interface CategoryBadgeProps {
  category: HookCategory;
}

const categoryStyles: Record<HookCategory, string> = {
  Formatting: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200',
  Notifications: 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200',
  Logging: 'bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-200',
  Validation: 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200',
  Security: 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200',
  Integration: 'bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-200',
  Productivity: 'bg-indigo-100 text-indigo-800 dark:bg-indigo-900 dark:text-indigo-200',
  Other: 'bg-neutral-100 text-neutral-800 dark:bg-neutral-700 dark:text-neutral-200',
};

export function CategoryBadge({ category }: CategoryBadgeProps) {
  return (
    <span
      className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${categoryStyles[category]}`}
    >
      {category}
    </span>
  );
}
