import type { Meta, StoryObj } from '@storybook/react';

import { TasksPage } from '../../pages/Tasks';
import { RecoilRoot } from 'recoil';
import { BrowserRouter } from 'react-router-dom';

const meta = {
  component: TasksPage,
  decorators: [
    (Story: any) => (
      <BrowserRouter>
      <RecoilRoot>
      <Story />
      </RecoilRoot>
      </BrowserRouter>
    ),
  ],
} satisfies Meta<typeof TasksPage>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};