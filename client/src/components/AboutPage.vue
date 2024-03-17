<template>
  <div>
    <NavBarElement />
    <div class="container-about">
      <h2 class="about-header">About application</h2>
      <div class="about-text">
        <ul>
          <li>
            <strong>Summarization:</strong> Summarize lengthy web content or
            data obtained through scraping into concise summaries. It employs
            both abstractive and extractive summarization techniques to generate
            condensed versions of articles, reviews, or forum threads. By
            providing users with summarized content, the application helps save
            time and improves readability, making it easier to digest large
            volumes of information.
          </li>
          <li>
            <strong>Sentiment Analysis:</strong> The application performs
            sentiment analysis on scraped data. It analyzes the overall
            sentiment or opinion expressed in the data, allowing users to gain
            insights into customer feedback, reviews, or social media comments.
            By offering sentiment analysis insights, the application enables
            users to better understand the sentiment behind the data, helping
            them make informed decisions and tailor their actions accordingly.
          </li>
          <li>
            <strong>Translation:</strong> The application offers translation
            capabilities, allowing users to translate content into another
            language. This feature enables users to overcome language barriers
            and access information in their preferred language.
          </li>
        </ul>
      </div>
      <FooterElement />
    </div>
  </div>
</template>

<style>
.container-about {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.about-header {
  margin-top: 1em;
  margin-bottom: 1em;
}

.about-text {
  width: 60%;
  text-align: justify;
}

.about-text li {
  margin-top: 1em;
}
</style>

<script>
import FooterElement from "./FooterElement.vue";
import NavBarElement from "./NavBarElement.vue";

export default {
  components: {
    FooterElement,
    NavBarElement,
  },
};
</script>
