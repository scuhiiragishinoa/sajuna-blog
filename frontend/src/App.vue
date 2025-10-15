<template>
    <div class="app">
      <header class="app-header">
        <h1>🎉 Sajuna Blog</h1>
        <p>Development Environment is Running!</p>
      </header>
      <main class="app-main">
        <div class="status-cards">
          <div class="status-card">
            <h3>Frontend</h3>
            <p>✅ Running on port 3000</p>
          </div>
          <div class="status-card">
            <h3>Backend API</h3>
            <p>✅ Running on port 8080</p>
          </div>
          <div class="status-card">
            <h3>MySQL</h3>
            <p>✅ Running on port 3306</p>
          </div>
          <div class="status-card">
            <h3>Redis</h3>
            <p>✅ Running on port 6379</p>
          </div>
        </div>
      </main>
    </div>
  </template>
  
  <script setup>
  //...
  </script>
  
  <style>
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }
  
  body {
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
  }
  
  .app {
    min-height: 100vh;
    color: white;
  }
  
  .app-header {
    text-align: center;
    padding: 4rem 2rem 2rem;
  }
  
  .app-header h1 {
    font-size: 3rem;
    margin-bottom: 1rem;
  }
  
  .app-header p {
    font-size: 1.2rem;
    opacity: 0.9;
  }
  
  .app-main {
    padding: 2rem;
    display: flex;
    justify-content: center;
  }
  
  .status-cards {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 1.5rem;
    max-width: 1200px;
    width: 100%;
  }
  
  .status-card {
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 12px;
    padding: 2rem;
    text-align: center;
    border: 1px solid rgba(255, 255, 255, 0.2);
  }
  
  .status-card h3 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
  }
  
  .status-card p {
    font-size: 1.1rem;
    opacity: 0.9;
  }
  
  @media (max-width: 768px) {
    .app-header h1 {
      font-size: 2rem;
    }
    
    .status-cards {
      grid-template-columns: 1fr;
    }
  }
  </style>