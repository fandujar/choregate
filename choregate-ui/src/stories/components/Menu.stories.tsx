import type { Meta, StoryObj } from '@storybook/react';

import { Menu } from '../../components/Menu';
import { BrowserRouter } from 'react-router-dom';

const meta = {
  component: Menu,
  decorators: [
    (Story: any) => (
      <BrowserRouter>
      <Story />
      </BrowserRouter>
    ),
  ],
} satisfies Meta<typeof Menu>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};