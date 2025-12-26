# Project Overview

## Overview
Star Wars Unlimited: What Would You Do (SWUWWYD) is a web application designed to help Star Wars Unlimited players improve their decision-making skills by creating and voting on hypothetical game scenarios.

## Target Audience
Star Wars Unlimited card game players who want to:
- Improve their strategic decision-making
- Learn from community insights
- Discuss optimal plays in specific board states
- Practice evaluating complex game situations

## Core Features
**User Authentication:**
- User registration and login
- Session management via Pocketbase

**Poll Creation:**
- Users can create new polls representing specific game scenarios
- Board state setup interface for defining:
    - Cards in play
    - Resources available
    - Game phase/turn information
    - Any other relevant game state details

**Poll Sharing:**
- Each published poll generates a unique URL with a GUID
- Shareable links allow community voting
- No authentication required to view polls, but required to vote on polls

**Voting System:**
- Users can vote on what action they would take in the presented scenario
- Vote tallying and results display

## Technology Stack
- **Language:** Golang
- **Frontend:** htmx with Golang templating
- **Backend:** Pocketbase
- **Authentication:** Pocketbase
- **Database:** Pocketbase
- **Testing:** `testify/assert` and `testify/require`