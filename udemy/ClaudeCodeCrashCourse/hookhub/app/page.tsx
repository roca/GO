import { Header } from '@/components/Header';
import { HookGrid } from '@/components/HookGrid';
import { hooks } from '@/lib/data/hooks';

export default function Home() {
  return (
    <div className="min-h-screen bg-gray-50 dark:bg-neutral-950">
      <Header />
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 md:py-12 lg:py-16">
        <HookGrid hooks={hooks} />
      </main>
    </div>
  );
}
