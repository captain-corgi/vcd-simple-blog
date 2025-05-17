import { Suspense, lazy } from 'react';
import { Routes, Route } from 'react-router-dom';
import { Layout } from '@/components/layout';
import { LoadingSpinner } from '@/components/ui/loading-spinner';

// Lazy-loaded micro frontends
const BlogModule = lazy(() => import('@/modules/blog'));
const AuthModule = lazy(() => import('@/modules/auth'));
const ProfileModule = lazy(() => import('@/modules/profile'));
const AdminModule = lazy(() => import('@/modules/admin'));

// Home page is part of the shell
import HomePage from '@/pages/home';
import NotFoundPage from '@/pages/not-found';

function App() {
  return (
    <Layout>
      <Suspense fallback={<LoadingSpinner />}>
        <Routes>
          <Route path="/" element={<HomePage />} />
          
          {/* Micro Frontend Routes */}
          <Route path="/blog/*" element={<BlogModule />} />
          <Route path="/auth/*" element={<AuthModule />} />
          <Route path="/profile/*" element={<ProfileModule />} />
          <Route path="/admin/*" element={<AdminModule />} />
          
          {/* 404 Route */}
          <Route path="*" element={<NotFoundPage />} />
        </Routes>
      </Suspense>
    </Layout>
  );
}

export default App;
