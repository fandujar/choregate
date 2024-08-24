import { BrowserRouter } from 'react-router-dom';

import { Router } from './Router';
import { AuthProvider } from './hooks/Auth';
import { RecoilRoot } from 'recoil';
import { Toaster } from '@/components/ui/sonner';

function App() {
  return (
    <div className='bg-slate-200 text-slate-950 h-full'>
    <RecoilRoot>
    <BrowserRouter>
    <AuthProvider>
      <Router />
      <Toaster />
    </AuthProvider>
    </BrowserRouter>
    </RecoilRoot>
    </div>
  )
}

export default App
