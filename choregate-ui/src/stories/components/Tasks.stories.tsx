import type { Meta, StoryObj } from '@storybook/react';

import { Tasks } from '../../components/Tasks';
import { RecoilRoot } from 'recoil';
import { BrowserRouter } from 'react-router-dom';

const meta = {
  component: Tasks,
  decorators: [
    (Story: any) => (
      <BrowserRouter>
      <RecoilRoot>
      <Story />
      </RecoilRoot>
      </BrowserRouter>
    ),
  ],
} satisfies Meta<typeof Tasks>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};