import type { Meta, StoryObj } from '@storybook/react';

import { Task } from '../../components/Task';
import { RecoilRoot } from 'recoil';

const meta = {
  component: Task,
  decorators: [
    (Story: any) => (
      <RecoilRoot>
        <Story />
      </RecoilRoot>
    ),
  ],
} satisfies Meta<typeof Task>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};