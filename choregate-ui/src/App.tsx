import { BrowserRouter } from 'react-router-dom';

import { Router } from './Router';
import { AuthProvider } from './hooks/Auth';
import { RecoilRoot } from 'recoil';

function App() {
  return (
    <div className='bg-slate-200 text-slate-950 h-full'>
    <RecoilRoot>
    <BrowserRouter>
    <AuthProvider>
      <Router />
    </AuthProvider>
    </BrowserRouter>
    </RecoilRoot>
    </div>
  )
}

export default App
