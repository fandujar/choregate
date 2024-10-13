import { BrowserRouter } from 'react-router-dom';

import { Router } from './Router';
import { AuthProvider } from './hooks/Auth';
import { RecoilRoot } from 'recoil';
import { Toaster } from '@/components/ui/sonner';
import {QueryClientProvider, QueryClient} from 'react-query'

function App() {
  const queryClient = new QueryClient()

  return (
    <div className='bg-gray-100 text-slate-950 h-full'>
    <RecoilRoot>
    <BrowserRouter>
    <AuthProvider>
      <QueryClientProvider client={queryClient}>
      <Router />
      <Toaster />
      </QueryClientProvider>
    </AuthProvider>
    </BrowserRouter>
    </RecoilRoot>
    </div>
  )
}

export default App
