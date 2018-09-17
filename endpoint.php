<?php
if (file_exists('cache.txt') && time() - filemtime('cache.txt') < 300) {
    header('Content-Type: application/json');
    readfile("cache.txt");
} else {
    //data collection goes here!
    //but i left it out for github to hide the api shizzle
  
    $jsonEftelOutput = $timeCache;
    $jsonMeta        = $localizedEnCache;
    $jsonParkMeta    = $timeCache;
    $jsonAppMeta     = $localizedEnCache;
    $jsonAppMetaNL   = $localizedNlCache;
    $jsonMaintenance = $timeCache->MaintenanceInfo;
    $rides           = $jsonEftelOutput->AttractionInfo;
    $maintenance     = $jsonMaintenance;
    
    $outRides   = new stdClass();
    $generalOut = new stdClass();
    
    //efteling api
    foreach ($rides as $ride) {
        $name                  = $ride->Id;
        $outRides->$name       = new stdClass();
        $outRides->$name->type = $ride->Type;
        
        //show
        if ($ride->Type == "Show") {
            if ($ride->ShowTimes != null) {
                $outRides->$name->showTimes = $ride->ShowTimes;
            } else {
                $outRides->$name->showTimes = array();
            }
            $outRides->$name->attractionType    = "SHOW";
        }
        
        //ride waiting time
        if ($ride->Type == "Attraction") {
            if ($ride->WaitingTime == null) {
                $ride->WaitingTime     = 0;
                $ride->StatePercentage = 0;
            }
            $outRides->$name->waitingTime       = $ride->WaitingTime;
            $outRides->$name->waitingPercentage = $ride->StatePercentage;
            $outRides->$name->attractionType    = "ATTRACTION";
            
            $outRides->$name->State = $ride->State;
            if ($ride->State == "inonderhoud")
                $outRides->$name->State = "maintenance";
            
        }
        
        if ($ride->Type == "Horeca") {
            $outRides->$name->openingTimes = $ride->OpeningHours;
            if ($ride->OpeningHours == null) $outRides->$name->openingTimes = array();
            $outRides->$name->attractionType    = "CATERING";
        }
        
        //real data
        if ($jsonMeta->rides->$name != null) {
            if ($jsonMeta->rides->$name->realname == null) {
                $jsonMeta->rides->$name->realname = $name;
            } else {
                $outRides->$name->realName = $jsonMeta->rides->$name->realname;
            }
        }
    }
    
    if (!isset($_GET["nometa"])) {
        //efteling app api
      foreach ($jsonAppMetaNL as $ride) {
            $name   = $ride->id;
            $data   = $ride->fields;
            $rideid = str_replace("-nl", "", $name);
            
            if ($outRides->$rideid != null) {
                if ($data->alternateid != null) {
                    $outRides->$rideid->alternateId = $data->alternateid;
                }
                
                if ($data->text != null) {
                    $outRides->$rideid->text_nl = $data->text;
                }
                
                if ($data->detail_text != null) {
                    $outRides->$rideid->detail_text_nl = strip_tags($data->detail_text);
                }
                
                if ($data->empire != null) {
                    $outRides->$rideid->area = $data->empire;
                }
                
                if ($data->properties != null) {
                    $outRides->$rideid->properties = $data->properties;
                }
                
                if ($data->showduration != null) {
                    if ($data->showduration != 0) {
                        $outRides->$rideid->showduration = $data->showduration;
                    }
                }
                
                if ($data->type != null) {
                    $outRides->$rideid->type = $data->type;
                }
                
                if ($data->latlon != null) {
                    $outRides->$rideid->location = $data->latlon;
                }
                
                $photos     = array();
                $photoFound = false;
                if ($data->image_detailview1 != null) {
                    $photoFound = true;
                    array_push($photos, "http://efteling.com" . $data->image_detailview1);
                }
                
                if ($data->image_detailview2 != null) {
                    array_push($photos, "http://efteling.com" . $data->image_detailview2);
                }
                
                if ($data->image_detailview3 != null) {
                    array_push($photos, "http://efteling.com" . $data->image_detailview3);
                }
                
                if ($data->image_detailview4 != null) {
                    array_push($photos, "http://efteling.com" . $data->image_detailview4);
                }
                
                if ($photoFound == true) {
                    $outRides->$rideid->photos = $photos;
                }
            }
        }
      
        foreach ($jsonAppMeta as $ride) {
            $name   = $ride->id;
            $data   = $ride->fields;
            $rideid = str_replace("-en", "", $name);
            
            if ($outRides->$rideid != null) {
                if ($data->alternateid != null) {
                    $outRides->$rideid->alternateId = $data->alternateid;
                }
                
                if ($data->text != null) {
                    $outRides->$rideid->text_en = $data->text;
                }
                
                if ($data->detail_text != null) {
                    $outRides->$rideid->detail_text_en = strip_tags($data->detail_text);
                }
                
                if ($data->empire != null) {
                    $outRides->$rideid->area = $data->empire;
                }
                
                if ($data->properties != null) {
                    $outRides->$rideid->properties = $data->properties;
                }
                
                if ($data->showduration != null) {
                    if ($data->showduration != 0) {
                        $outRides->$rideid->showduration = $data->showduration;
                    }
                }
                
                if ($data->type != null) {
                    $outRides->$rideid->type = $data->type;
                }
                
                if ($data->latlon != null) {
                    $outRides->$rideid->location = $data->latlon;
                }
                
                $photos     = array();
                $photoFound = false;
                if ($data->image_detailview1 != null) {
                    $photoFound = true;
                    array_push($photos, "http://efteling.com" . $data->image_detailview1);
                }
                
                if ($data->image_detailview2 != null) {
                    array_push($photos, "http://efteling.com" . $data->image_detailview2);
                }
                
                if ($data->image_detailview3 != null) {
                    array_push($photos, "http://efteling.com" . $data->image_detailview3);
                }
                
                if ($data->image_detailview4 != null) {
                    array_push($photos, "http://efteling.com" . $data->image_detailview4);
                }
                
                if ($photoFound == true) {
                    $outRides->$rideid->photos = $photos;
                }
            }
        }
    }
    
    $maintenanceOutput = new stdClass();
    
    //efteling maintenance
    foreach ($maintenance as $key) {
        $renderedObject                = new stdClass();
        $renderedObject->from          = $key->DateFrom;
        $renderedObject->end           = $key->DateTo;
        $ride                          = $key->AttractionId;
        $renderedObject->openInWeekend = $key->OpenInWeekend;
        $outRides->$ride->maintenance  = $renderedObject;
        $maintenanceOutput->$ride      = $renderedObject;
    }
    
    $generalOut->lastUpdate            = date("h:i:s");
    $generalOut->maintenanceInfo       = $maintenanceOutput;
    $generalOut->park                  = new stdClass();
    $generalOut->park->openFrom        = $jsonParkMeta->OpeningHours->HourFrom;
    $generalOut->park->openUntil       = $jsonParkMeta->OpeningHours->HourTo;
    $generalOut->park->busyIndication  = $jsonParkMeta->OpeningHours->BusyIndication;
    $generalOut->rides                 = $outRides;
    
    header('Content-Type: application/json');
    $o = json_encode($generalOut);
    echo $o;

    ob_start();
    $fp = fopen('cache.txt', "w");
    fwrite($fp, $o);
    fclose($fp);
    ob_end_flush();
}
?>
