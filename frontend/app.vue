<!-- Reactのreturnの部分 -->
<template>
  <div>
    <nav class="navbar">
      <div class="nav-container">
        <h1>EDINET Analyzer</h1>
      </div>
    </nav>
    
    <main class="main-content">
      <NuxtPage />
    </main>
  </div>
</template>

<script>
  // JSの部分
</script>

<!-- CSSの部分 -->
<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background: #f3f4f6;
}

.navbar {
  background: white;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  padding: 1rem 0;
  margin-bottom: 2rem;
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
}

.navbar h1 {
  color: #2563eb;
  font-size: 1.5rem;
}

.main-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
}
</style>