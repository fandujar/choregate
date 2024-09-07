import type { Meta, StoryObj } from '@storybook/react';

import { Admin } from '../../components/Admin';

const meta = {
  component: Admin,
} satisfies Meta<typeof Admin>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};