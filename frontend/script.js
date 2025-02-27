document.addEventListener('DOMContentLoaded', () => {
    const postContent = document.getElementById('postContent');
    const submitButton = document.getElementById('submitPost');
    const postsContainer = document.getElementById('posts');
    const latestResponse = document.getElementById('latestResponse');

    latestResponse.classList.add('empty');
    latestResponse.textContent = 'AI response will appear here...';

    async function fetchPosts() {
        try {
            const response = await fetch('http://localhost:8080/posts');
            const posts = await response.json();
            displayPosts(posts);
        } catch (error) {
            console.error('Error fetching posts:', error);
        }
    }

    function displayPosts(posts) {

        const sortedPosts = [...posts].sort((a, b) => 
            new Date(b.created_at) - new Date(a.created_at)
        );

        postsContainer.innerHTML = sortedPosts.map(post => `
            <div class="post">
                <div class="post-content">${post.content}</div>
                <div class="post-metadata">
                    <span class="sentiment-tag">Sentiment: ${post.sentiment}</span>
                    <span>${new Date(post.created_at).toLocaleDateString()}</span>
                </div>
                <div class="post-response">
                    Response: ${post.response}
                </div>
            </div>
        `).join('');
    }

    submitButton.addEventListener('click', async () => {
        const content = postContent.value.trim();
        if (!content) return;

        submitButton.disabled = true;
        submitButton.textContent = 'Analyzing...';

        try {
            const response = await fetch('http://localhost:8080/posts', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ content }),
            });

            if (response.ok) {
                const result = await response.json();

                latestResponse.classList.remove('empty');
                latestResponse.innerHTML = `
                    <div class="post-content">${result.content}</div>
                    <div class="post-metadata">
                        <span class="sentiment-tag">Sentiment: ${result.sentiment}</span>
                    </div>
                    <div class="post-response">
                        Response: ${result.response}
                    </div>
                `;
                
                postContent.value = '';
                fetchPosts(); 
            }
        } catch (error) {
            console.error('Error creating post:', error);
        } finally {
            submitButton.disabled = false;
            submitButton.textContent = 'Analyze & Submit';
        }
    });

    fetchPosts();
}); 