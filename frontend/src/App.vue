<template>
  <!-- Header -->
  <Header/>

  <!-- Main Header -->
  <div class="main-header">

    <!-- Row 1 -->
  <Row1/>

    <!-- Row 2 -->
  <Row2/>

    <!-- Main Body -->
    <table class="main-body">
      <tbody>
        <tr>
          <!-- Sidebar -->
          <td class="sidebar" style="padding-right: 6px">
            <table class="sidebar-table">
              <tbody>
                <tr><td><img src="/images/secure.gif" alt="secure"></td></tr>
                <tr class="quicklink"><td>Quick Links</td></tr>
                <tr class="search">
                  <td>
                    BongThom-ID:
                    <input type="text" name="ID" size="5">
                    <input type="submit" value=">" name="btnSearch">
                  </td>
                </tr>
              </tbody>
              <QuickLinks/>
              <tbody>
                <tr><td style="height: 7px;"></td></tr>
                <tr class="joblist"><td>Job Quick Lists</td></tr>
                <tr><td style="height: 7px;"></td></tr>
                <tr class="schedule"><td>Special Schedule</td></tr>
              </tbody>
              <SpecialSchedule/>
              <tbody>
                <tr><td style="height: 7px;"></td></tr>
                <tr class="organization"><td>Organization Type</td></tr>
              </tbody>
              <OrgType />
              <tbody>
                <tr><td style="height: 7px;"></td></tr>
                <tr class="other-announcements"><td>Other Announcements</td></tr>
              </tbody>
              <OtherAnn/>
              <tbody>
                <tr><td style="height: 7px;"></td></tr>
                <tr class="leftinfo" id="jobcategory">
                  <td>»<a href="#">Khmer / English / French <br> MP3 Phrasebook</a></td>
                </tr>
                <tr><td style="height: 12px"></td></tr>
                <tr class="rss-files" style="color: #800000;"><td>RSS Files</td></tr>
              </tbody>
              <RssFiles/>
              <tbody>
                <tr class="rss-info" style="text-align: center" id="click-here">
                  <td><b><a href="#" style="color: purple; text-decoration: underline;">Click here</a> to download an RSS<br>Browser</b></td>
                </tr>
                <tr><td style="height: 12px;"></td></tr>
                <tr class="leftinfo">
                  <td class="sidebar-note">* NOTE: The numbers in brackets above () are the number of announcements containing job positions in that category</td>
                </tr>
              </tbody>
            </table>
          </td>

          <!-- Main Content -->
          <td class="main-content" style="vertical-align: top;">
            <table style="width: 100%; border-collapse: collapse;">
              <tbody>
                <tr>
                  <!-- Jobs Column -->
                  <td style="width: 50%; vertical-align: top; border-right: 2px dashed black;">
                    <table style="width: 100%; border-collapse: collapse;">
                      <tbody>
                        <tr><td class="header-title-body">Quick View Latest Jobs <span class="random">(<img src="/images/new.gif" alt="New">= New)</span></td></tr>
                        <tr class="search1">
                          <td>
                            <strong>Search Job Title</strong>
                            <input type="text" name="ID" size="6">
                            <input type="submit" value=">" name="btnSearch">
                          </td>
                        </tr>
                        <tr>
                          <td class="jobs-background">
                            <ul><JobList/></ul>
                          </td>
                        </tr>
                        <tr><td class="view-alljobs"><a href="#">Quick View All Active Jobs</a></td></tr>
                      </tbody>
                    </table>
                  </td>

                  <!-- Classifieds Column -->
                  <td style="width: 50%; vertical-align: top;">
                    <table style="width: 100%; border-collapse: collapse">
                      <tbody>
                        <tr><td class="header-title-right">Quick View Latest Classifieds</td></tr>
                        <tr class="search2">
                          <td>
                            <strong>Search Classifieds Title</strong>
                            <input type="text" name="ID" size="6">
                            <input type="submit" value=">" name="btnSearch">
                          </td>
                        </tr>
                        <tr>
                          <td class="classified-background">
                            <ul><ClassifiedList /></ul>
                          </td>
                        </tr>
                        <tr><td class="view-allclassifieds"><a href="#">Quick View All Active Classifieds</a></td></tr>
                      </tbody>
                    </table>
                  </td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>

    <hr class="steav-hr">

    <!-- Footer -->
   <Footer/>

  </div>
</template>

<script setup>
import QuickLinks from './components/QuickLinks.vue'
import SpecialSchedule from './components/SpecialSchedule.vue'
import OrgType from './components/OrgType.vue'
import OtherAnn from './components/OtherAnn.vue'
import RssFiles from './components/RssFiles.vue'
import JobList from './components/JobList.vue'
import ClassifiedList from './components/ClassifiedList.vue'
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
import Row1 from "./components/Row1.vue";
import Row2 from "./components/Row2.vue";
</script>