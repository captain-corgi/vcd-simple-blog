# Micro Frontend Architecture

The VCD Simple Blog application implements a Micro Frontend architecture to enable independent development, deployment, and scaling of frontend features.

## Overview

Micro Frontend architecture extends the concept of microservices to the frontend, breaking down a monolithic frontend into smaller, more manageable pieces that can be developed, tested, and deployed independently.

## Implementation Approach

We use a **runtime integration** approach with a shell application that dynamically loads feature modules at runtime.

### Shell Application

The shell application is responsible for:
- Providing the application shell (header, footer, navigation)
- Managing routing and navigation between features
- Loading and rendering feature modules
- Managing global state and authentication
- Providing shared services and utilities

### Feature Modules

Feature modules are independent applications that:
- Focus on specific business domains or user journeys
- Can be developed and deployed independently
- Share common UI components and utilities
- Communicate with the shell and other modules through well-defined interfaces

## Technical Implementation

### Module Federation

We use Vite's module federation capabilities to enable runtime loading of feature modules:

```javascript
// vite.config.ts for feature modules
export default defineConfig({
  build: {
    lib: {
      entry: path.resolve(__dirname, 'src/index.ts'),
      name: 'BlogModule',
      formats: ['es'],
      fileName: 'blog-module',
    },
    rollupOptions: {
      external: ['react', 'react-dom', 'react-router-dom', '@tanstack/react-query'],
      output: {
        globals: {
          react: 'React',
          'react-dom': 'ReactDOM',
          'react-router-dom': 'ReactRouterDOM',
          '@tanstack/react-query': 'ReactQuery',
        },
      },
    },
  },
});
```

### Dynamic Loading

The shell application dynamically loads feature modules using React's lazy loading:

```jsx
// App.tsx in shell application
import { Suspense, lazy } from 'react';
import { Routes, Route } from 'react-router-dom';
import { Layout } from '@/components/layout';
import { LoadingSpinner } from '@/components/ui/loading-spinner';

// Lazy-loaded micro frontends
const BlogModule = lazy(() => import('@/modules/blog'));
const AuthModule = lazy(() => import('@/modules/auth'));
const ProfileModule = lazy(() => import('@/modules/profile'));
const AdminModule = lazy(() => import('@/modules/admin'));

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
```

### Shared Components and Utilities

Common UI components and utilities are shared through internal packages:

```
packages/
  ui/             # Shared UI components
  utils/          # Shared utilities
  hooks/          # Shared React hooks
  api-client/     # Shared API client
```

These packages are consumed by both the shell and feature modules using pnpm workspaces.

## Communication Between Modules

### Props and Context

The shell passes data and callbacks to feature modules through props and React Context:

```jsx
// Shell application
<BlogModule 
  user={currentUser} 
  onBlogCreated={handleBlogCreated} 
/>
```

### Event Bus

For more complex communication, we use a simple event bus:

```typescript
// Event bus implementation
export const eventBus = {
  on(event, callback) {
    document.addEventListener(event, (e) => callback(e.detail));
  },
  dispatch(event, data) {
    document.dispatchEvent(new CustomEvent(event, { detail: data }));
  },
  remove(event, callback) {
    document.removeEventListener(event, callback);
  },
};
```

### Shared State

For global state that needs to be shared across modules, we use React Query for server state and Context API for client state.

## Benefits of Micro Frontend Architecture

1. **Independent Development**: Teams can work on different features without stepping on each other's toes.
2. **Incremental Upgrades**: Features can be upgraded independently without affecting the entire application.
3. **Technology Flexibility**: Different teams can use different technologies if needed (though we maintain consistency).
4. **Scalable Development**: More teams can work on the application simultaneously.
5. **Focused Codebases**: Each feature module has a smaller, more focused codebase.

## Challenges and Mitigations

1. **Consistency**: We maintain consistency through shared UI components and design guidelines.
2. **Performance**: We optimize bundle sizes and use code splitting to minimize performance impact.
3. **Complexity**: We document integration points and communication patterns clearly.
4. **Testing**: We have end-to-end tests that verify integration between modules.
