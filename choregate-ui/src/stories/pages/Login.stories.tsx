import type { Meta, StoryObj } from '@storybook/react';

import { LoginPage } from '../../pages/Login';

const meta = {
  component: LoginPage,
  decorators: [
    (Story: any) => (
      <Story />
    ),
  ],
} satisfies Meta<typeof LoginPage>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};