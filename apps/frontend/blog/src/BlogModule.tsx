import { Routes, Route } from 'react-router-dom';
import BlogListPage from './pages/BlogListPage';
import BlogDetailPage from './pages/BlogDetailPage';
import CreateBlogPage from './pages/CreateBlogPage';
import EditBlogPage from './pages/EditBlogPage';

function BlogModule() {
  return (
    <Routes>
      <Route path="/" element={<BlogListPage />} />
      <Route path="/:id" element={<BlogDetailPage />} />
      <Route path="/create" element={<CreateBlogPage />} />
      <Route path="/:id/edit" element={<EditBlogPage />} />
    </Routes>
  );
}

export default BlogModule;
