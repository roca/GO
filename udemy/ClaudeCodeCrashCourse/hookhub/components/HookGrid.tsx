import { Hook } from '@/lib/types/hook';
import { HookCard } from './HookCard';

interface HookGridProps {
  hooks: Hook[];
}

export function HookGrid({ hooks }: HookGridProps) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {hooks.map((hook) => (
        <HookCard key={hook.id} hook={hook} />
      ))}
    </div>
  );
}
