// Register API Key here for more requests & APIs: https://newsapi.org
const API_KEY = "75d5292ba7504820a7eda8d0f2c74199";

export async function getNews(page = 1, pageSize = 20) {
  const sources = "bbc-news,cbc-news,nbc-news,fox-news,mtv-news";
 // const response = await fetch(
 //   `https://newsapi.org/v2/top-headlines?sources=${sources}=&page=${page}&pageSize=${pageSize}&apiKey=${API_KEY}`
  const response = await fetch(
    `https://newsapi.org/v2/top-headlines?country=br&apiKey=c800640326d24b22ab17a1ab64f620f8`
  );
  const jsonData = await response.json();
  return jsonData.articles || [];
}
