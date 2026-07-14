export type HookEvent =
  | 'PreToolUse'
  | 'PermissionRequest'
  | 'PostToolUse'
  | 'ToolError'
  | 'SessionStart'
  | 'SessionEnd'
  | 'MessageSent'
  | 'ResponseReceived';

export type HookCategory =
  | 'Formatting'
  | 'Notifications'
  | 'Logging'
  | 'Validation'
  | 'Security'
  | 'Integration'
  | 'Productivity'
  | 'Other';

export interface Hook {
  id: string;
  name: string;
  description: string;
  category: HookCategory;
  repoUrl: string;
  author?: string;
  tags?: string[];
  hookEvent: HookEvent;
  createdAt?: string;
}
