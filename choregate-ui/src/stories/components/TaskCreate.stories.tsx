import type { Meta, StoryObj } from '@storybook/react';

import { TaskCreate } from '../../components/TaskCreate';
import { RecoilRoot } from 'recoil';

const meta = {
  component: TaskCreate,
  decorators: [
    (Story: any) => (
      <RecoilRoot>
        <Story />
      </RecoilRoot>
    ),
  ],
} satisfies Meta<typeof TaskCreate>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};